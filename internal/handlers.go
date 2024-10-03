package internal

import (
	"encoding/json"
	"net/http"
	"strconv"
)

// Получение всех баннеров
func GetBannersHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := DB.Query("SELECT * FROM banners WHERE is_active = TRUE")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var banners []map[string]interface{}
	for rows.Next() {
		var id int
		var jsonData string
		var featureID int
		var isActive bool
		if err := rows.Scan(&id, &jsonData, &featureID, &isActive); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		banners = append(banners, map[string]interface{}{
			"id":         id,
			"json_data":  jsonData,
			"feature_id": featureID,
			"is_active":  isActive,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(banners)
}

// Создание нового баннера
func CreateBannerHandler(w http.ResponseWriter, r *http.Request) {
	var banner map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&banner); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Пример добавления баннера в базу данных
	_, err := DB.Exec("INSERT INTO banners (json_data, feature_id, is_active) VALUES ($1, $2, $3)",
		banner["json_data"], banner["feature_id"], banner["is_active"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// Обновление баннера
func UpdateBannerHandler(w http.ResponseWriter, r *http.Request) {
	// Извлечение ID баннера из URL
	idStr := r.URL.Path[len("/banners/update/"):] // Получаем часть пути после "/banners/update/"

	// Преобразование ID в целое число
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	// Получение новых данных из тела запроса
	var banner map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&banner); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Выполнение запроса на обновление
	_, err = DB.Exec("UPDATE banners SET json_data = $1, feature_id = $2, is_active = $3 WHERE id = $4",
		banner["json_data"], banner["feature_id"], banner["is_active"], id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Ответ клиенту о успешном обновлении
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Banner updated successfully"})
}

// Удаление баннера
func DeleteBannerHandler(w http.ResponseWriter, r *http.Request) {
	// Извлечение ID баннера из URL
	idStr := r.URL.Path[len("/banners/delete/"):]

	// Преобразование ID в целое число
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	// Выполнение запроса на удаление
	_, err = DB.Exec("DELETE FROM banners WHERE id = $1", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Ответ клиенту о успешном удалении
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Banner deleted successfully"})
}
