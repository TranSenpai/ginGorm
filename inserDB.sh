
code=2180606028

for i in {1..100}
do
  $(( code += i ))
  echo $code
done

curl --location 'http://localhost:8080/api/v1/contract/' \
--header 'Content-Type: application/json' \
--data-raw '{          
	"StudentCode"           : "2180606025",
	"FullName"              : "Tran Sy Chuong",
	"Email"                 : "chuongtran975@gmail.com",
	"Sign"                  : "abc123", 
	"Phone"                 : "0123456789",
	"Gender"                : 1,
	"DOB"                   : "2003-12-24T15:04:05Z",
	"Address"               : "abc123",
	"Avatar"                : "qweqweqweqweqweqweqwe",
	"IsActive"              : true,
	"RoomID"                : "B0102",
	"NotificationChannels"  : "sms"
}'

done
	# // Delay để đảm bảo server đã chạy trước khi gửi request
	# time.Sleep(2 * time.Second)

	# // Tạo random student code
	# rand.Seed(time.Now().UnixNano())
	# num := rand.Intn(9999999999-2180606090) + 2180606090

	# url := "http://localhost:8080/api/v1/contract/"
	# method := "POST"

	# payload := strings.NewReader(fmt.Sprintf(`{
	# "StudentCode"           : "%d",
	# "FullName"              : "Tran Sy Chuong",
	# "Email"                 : "chuongtran975@gmail.com",
	# "Sign"                  : "abc123",
	# "Phone"                 : "0123456789",
	# "Gender"                : 1,
	# "DOB"                   : "2003-12-24T15:04:05Z",
	# "Address"               : "abc123",
	# "Avatar"                : "0AAAAASUVORK5CYII=",
	# "IsActive"              : true,
	# "RoomID"                : "B0102",
	# "NotificationChannels" : "sms"
	# }`, num))

	# client := &http.Client{}
	# req, err := http.NewRequest(method, url, payload)
	# if err != nil {
	# 	fmt.Println("Error creating request:", err)
	# 	return
	# }
	# req.Header.Add("Content-Type", "application/json")

	# res, err := client.Do(req)
	# if err != nil {
	# 	fmt.Println("Error sending request:", err)
	# 	return
	# }
	# defer res.Body.Close()

	# body, err := io.ReadAll(res.Body)
	# if err != nil {
	# 	fmt.Println("Error reading response:", err)
	# 	return
	# }
	# fmt.Println("Response:", string(body))