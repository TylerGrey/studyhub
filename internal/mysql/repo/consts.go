package repo

const defaultListLimit = int32(20)
const defaultListCursor = "CONCAT(UNIX_TIMESTAMP(created_at), LPAD(id, 10, '0'))"
const defaultListOrderField = "CREATED_AT"
const defaultListOrderDirection = "DESC"
