compile_backend:
	cd backend && go build

run_frontend:
	cd frontend && npm run dev

run_backend:
	cd backend && ./lazydependency

run:
	make compile_backend
	make -j 2 run_frontend run_backend
