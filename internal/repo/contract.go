package repo

var ContractRepo *contractRepo

func GetInstanceContract() IContractRepo {
	if ContractRepo == nil {
		ContractRepo = &contractRepo{db: dbConn}
	}
	return ContractRepo
}
