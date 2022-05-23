BTL (BTL 13): Ứng dụng TLS để xây dựng ứng dụng gửi tin nhắn giữa nhiều người dùng (mô phỏng các ứng dụng nhắn tin).
Cách khởi động chương trình : 
. Chúng ta cần phải tải ứng dụng Visual Studio Code để có thể thực hiện các đoạn mã ( code ) , khi tải xong chúng ta tiến hành tạo 2 chương trình có tên "Server"
(máy chủ) , Client (khách ) , trong 2 chương trình ta cần thực hiện các dòng code cần thiết liên quan tới giao thức TLS , lúc hoàn thành chúng ta có thể chạy bằng 
Command Prompt , chúng ta phải mở Command Prompt của 2 thư mục chúng ta đã lưu vào máy là "Server " , "Client " để chạy chương trình 
. Có 2 cách chạy chương trình đó là chúng ta có thể chạy trực tiếp ở Visual Studio Code bằng cách mở New Terminal và sử dụng các câu lệnh như "cd sv.go " ... , 
cách thứ 2 chúng ta có thể mở trực tiếp file ta đã lưu ở máy và mở Command Prompt của 2 thư mục đó , khi mở đã lên chúng ta sử dụng câu lệnh "go run sv.go" để chạy
chương trình máy chủ (Server ) , sử dụng tiếp câu lệnh " go run cl.go" của chương trình khách (Client ) . Khi yêu cầu ở máy chủ hiện lên , chúng ta ấn nút "Allow" để chấp nhận
, còn bên phía khách (Client ) chúng ta điền tên người dùng để thực hiện kết nối tới Server . 
