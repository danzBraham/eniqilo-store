package repositories

import (
	"context"
	"strconv"
	"time"

	"github.com/danzBraham/eniqilo-store/internal/entities/checkoutentity"
	"github.com/danzBraham/eniqilo-store/internal/errors/checkouterror"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/oklog/ulid/v2"
)

type CheckoutRepository interface {
	CreateCheckoutProduct(ctx context.Context, transaction *checkoutentity.Transaction, checkoutProduct *checkoutentity.CheckoutProductRequest) error
	GetCheckoutHistories(ctx context.Context, params *checkoutentity.CheckoutHistoryQueryParams) ([]*checkoutentity.GetCheckoutHistoryResponse, error)
}

type CheckoutRepositoryImpl struct {
	DB *pgxpool.Pool
}

func NewCheckoutRepository(db *pgxpool.Pool) CheckoutRepository {
	return &CheckoutRepositoryImpl{DB: db}
}

func (r CheckoutRepositoryImpl) CreateCheckoutProduct(ctx context.Context, transaction *checkoutentity.Transaction, checkoutProduct *checkoutentity.CheckoutProductRequest) error {
	tx, err := r.DB.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	queryInsertTransaction := `INSERT INTO transactions (id, customer_id) VALUES ($1, $2)`
	_, err = tx.Exec(ctx, queryInsertTransaction, &transaction.ID, &transaction.CustomerID)
	if err != nil {
		return err
	}

	totalPriceTransaction := 0

	for _, pd := range checkoutProduct.ProductDetails {
		totalPriceProduct := 0

		var isProductExists bool
		queryIsProductExists := `SELECT EXISTS(SELECT 1 FROM products WHERE id = $1 AND is_deleted = false)`
		err := tx.QueryRow(ctx, queryIsProductExists, pd.ProductID).Scan(&isProductExists)
		if err != nil {
			return err
		}
		if !isProductExists {
			return checkouterror.ErrProductIDNotFound
		}

		var isProductAvailable bool
		queryIsProductAvailable := `SELECT is_available FROM products WHERE id = $1`
		err = tx.QueryRow(ctx, queryIsProductAvailable, pd.ProductID).Scan(&isProductAvailable)
		if err != nil {
			return err
		}
		if !isProductAvailable {
			return checkouterror.ErrOneOfProductNotAvailable
		}

		var stock, price int
		queryGetProductStockAndPrice := `SELECT stock, price FROM products WHERE id = $1`
		err = tx.QueryRow(ctx, queryGetProductStockAndPrice, pd.ProductID).Scan(&stock, &price)
		if err != nil {
			return err
		}
		if stock < pd.Quantity {
			return checkouterror.ErrOneOfProductStockNotEnough
		}

		queryUpdateProductStock := `UPDATE products SET stock = stock - $1 WHERE id = $2`
		_, err = tx.Exec(ctx, queryUpdateProductStock, pd.Quantity, pd.ProductID)
		if err != nil {
			return err
		}

		totalPriceProduct += price * pd.Quantity

		queryInsertCheckout := `
			INSERT INTO
				checkouts (id, transaction_id, product_id, quantity, total_price)
			VALUES
				($1, $2, $3, $4, $5)
		`
		_, err = tx.Exec(ctx, queryInsertCheckout,
			ulid.Make().String(),
			&transaction.ID,
			pd.ProductID,
			pd.Quantity,
			totalPriceProduct,
		)
		if err != nil {
			return err
		}

		totalPriceTransaction += totalPriceProduct
	}

	if checkoutProduct.Paid < totalPriceTransaction {
		return checkouterror.ErrPaidNotEnough
	}

	if checkoutProduct.Change != (checkoutProduct.Paid - totalPriceTransaction) {
		return checkouterror.ErrChangeNotRight
	}

	queryUpdateTransaction := `
		UPDATE
			transactions
		SET
			total_price = $1,
			paid = $2,
			change = $3
		WHERE
			id = $4
	`
	_, err = tx.Exec(ctx, queryUpdateTransaction,
		totalPriceTransaction,
		&checkoutProduct.Paid,
		&checkoutProduct.Change,
		&transaction.ID,
	)
	if err != nil {
		return err
	}

	if err := tx.Commit(ctx); err != nil {
		return err
	}

	return nil
}

func (r *CheckoutRepositoryImpl) GetCheckoutHistories(ctx context.Context, params *checkoutentity.CheckoutHistoryQueryParams) ([]*checkoutentity.GetCheckoutHistoryResponse, error) {
	query := `
		SELECT
			t.id, t.customer_id, t.paid, t.change, t.created_at,
			c.product_id, c.quantity
		FROM
			transactions t
		INNER JOIN
			checkouts c ON c.transaction_id = t.id
		WHERE
			is_deleted = false
	`
	args := []interface{}{}
	argID := 1

	if params.CustomerID != "" {
		query += ` AND customer_id = $` + strconv.Itoa(argID)
		args = append(args, params.CustomerID)
	}

	switch params.CreatedAt {
	case "asc":
		query += ` ORDER BY created_at ASC`
	case "desc":
		query += ` ORDER BY created_at DESC`
	}

	query += ` LIMIT $` + strconv.Itoa(argID) + ` OFFSET $` + strconv.Itoa(argID+1)
	args = append(args, params.Limit, params.Offset)

	rows, err := r.DB.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	historyMaps := make(map[string]*checkoutentity.GetCheckoutHistoryResponse)
	for rows.Next() {
		var (
			transactionID, customerID, productID string
			paid, change, quantity               int
			createdAt                            time.Time
		)
		err := rows.Scan(
			&transactionID,
			&customerID,
			&paid,
			&change,
			&createdAt,
			&productID,
			&quantity,
		)
		if err != nil {
			return nil, err
		}

		if _, exists := historyMaps[transactionID]; !exists {
			historyMaps[transactionID] = &checkoutentity.GetCheckoutHistoryResponse{
				TransactionID:  transactionID,
				CustomerID:     customerID,
				ProductDetails: []*checkoutentity.ProductDetails{},
				Paid:           paid,
				Change:         change,
				CreatedAt:      createdAt.Format(time.RFC3339),
			}
		}

		history := historyMaps[transactionID]
		history.ProductDetails = append(history.ProductDetails, &checkoutentity.ProductDetails{
			ProductID: productID,
			Quantity:  quantity,
		})
	}

	histories := make([]*checkoutentity.GetCheckoutHistoryResponse, 0, len(historyMaps))
	for _, history := range historyMaps {
		histories = append(histories, history)
	}

	return histories, nil
}
