package controllers

import "net/http"

func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Criando usuário"))
}
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando usuários"))
}
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscar um usuário"))
}
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizando um usuário"))
}
func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletando um usuário"))
}
