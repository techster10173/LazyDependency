compile_backend:
	cd backend && go build -o out/lazydependency

run_frontend:
	cd frontend && npm run dev

run_backend_dev:
	cd backend && go run server.go

run:
	make compile_backend
	make -j 2 run_frontend run_backend
