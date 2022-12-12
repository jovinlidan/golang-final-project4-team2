package transaction_domain

var TransactionDomain transactionDomainRepo = &transactionDomain{}

const (
	queryCreateTransaction = `INSERT INTO transactions (product_id, quantity) VALUES($1, $2)
	RETURNING total_price, quantity, product_title`
	queryGetMyTransactions = `SELECT t.id AS t_id, t.product_id AS t_product_id, t.user_id AS t_user_id, 
    quantity, total_price,p.id AS p_id, p.title, p.price, p.stock, p.category_id AS p_category_id, 
    p.created_at AS p_created_at, p.updated_at AS p_updated_at from transactions t 
    left join products p on t.product_id = p.id ORDER BY t.id`
)
