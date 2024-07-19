
# GoWeather

API ini menyediakan informasi cuaca terkini berdasarkan nama kota yang diberikan. 
Data cuaca diperoleh dari layanan OpenWeatherMap.

Setup and Configuration
-----------------------
1. Clone repository ini ke direktori lokal Anda.
2. Dapatkan API Key OpenWeatherMap:
   - Daftar di OpenWeatherMap (https://home.openweathermap.org/users/sign_up) dan dapatkan API key.
3. Buat file `.env` di root proyek dan tambahkan API key OpenWeatherMap Anda:
   ```plaintext
   OPENWEATHERMAP_API_KEY=your_openweathermap_api_key
4. Inisialisasi proyek:
   ```plaintext
   go mod tidy
5. Jalankan proyek
   ```plaintext
   go run main.go
6. Gunakan browser atau tools seperti Postman untuk mengakses endpoint berikut
    ```plaintext
   GET http://localhost:8080/weather/current?city=Jakarta


Dependencies
-----------------------
- Golang 1.19
- Gin Framework v1.7.7
- godotenv v1.3.0




## Endpoint

| URL                  | Method | Status | Request                                    | Response                                      |
|----------------------|--------|--------|--------------------------------------------|-----------------------------------------------|
| `/weather/current`   | GET    | 200    | `city` (query parameter)                   | ```json { "weather": { "name": "City Name", "main": { "temp": 25.0, "pressure": 1012, "humidity": 60 }, "weather": [ { "description": "clear sky" } ] } } ``` |
|                      |        | 400    | -                                          | ```json { "error": "City is required" } ```  |
|                      |        | 500    | -                                          | ```json { "error": "error message" } ```     |



