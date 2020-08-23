package models

type ByUpdatedAt []Job

func (u ByUpdatedAt) Len() int {
	return len(u)
}

func (u ByUpdatedAt) Swap(i, j int) {
	u[i], u[j] = u[j], u[i]
}

func (u ByUpdatedAt) Less(i, j int) bool {
	return u[i].UpdatedAt.After(u[j].UpdatedAt)
}
