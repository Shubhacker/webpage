package postgres

import (
	"context"
	"log"
)

func AuthRoleForUser(UserName string) string {
	log.Println("AuthRoleForUser()")
	if pool == nil {
		pool = GetPool()
	}
	var UserRole string
	QueryString := `select ur.user_role from user_table ut 
inner join user_role ur on ur.role_id = ut.user_role 
where ut.user_name = $1`
	err := pool.QueryRow(context.Background(), QueryString, UserName).Scan(&UserRole)
	if err != nil {
		log.Println("Error in AuthRoleForUser()")
	}
	return UserRole
}

func FetchSecretKey() string {
	log.Println("FetchSecretKey()")
	if pool == nil {
		pool = GetPool()
	}
	var SecretKey string
	QueryString := `select description from code where value = 'Secret Key'`
	err := pool.QueryRow(context.Background(), QueryString).Scan(&SecretKey)
	if err != nil {
		log.Println("Error in FetchSecretKey()")
	}
	return SecretKey
}
