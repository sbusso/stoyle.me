docker run --name stoyle -p 6379:6379 -d redis

cd frontend
npm install
npm run build
cd ..

cd backend
go build -o stoyle main.go
./stoyle