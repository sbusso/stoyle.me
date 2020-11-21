mkdir data
docker run -p 6379:6379 -d -v $PWD/data:/data --name stoyle redis:latest

cd frontend
npm install
npm run build
cd ..

cd backend
go build -o stoyle main.go
./stoyle