package repositories

import (
	"context"

	"github.com/danzBraham/eniqilo-store/internal/entities/checkoutentity"
	"github.com/danzBraham/eniqilo-store/internal/errors/checkouterror"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/oklog/ulid/v2"
)

type CheckoutRepository interface {
	CreateCheckoutProduct(ctx context.Context, transaction *checkoutentity.Transaction, checkoutProduct *checkoutentity.CheckoutProductRequest) error
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
				($1, $2, $3, $4 $5)
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
			total_price = $1
			paid = $2
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
