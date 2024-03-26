CONCURRENCY_TYPE=""
go-concurrency:
	@if [ $(CONCURRENCY_TYPE) = "CONCURRENT" ]; then\
		go run concurrency/main.go concurrent;\
	elif [ $(CONCURRENCY_TYPE) = "PIPELINE" ]; then\
		go run concurrency/main.go pipeline;\
	else\
		echo "🚨 You have to specify which package you want to run in the concurrent module 🚨\n\n💡Example: make go-concurrent CONCURRENCY_TYPE=CONCURRENT \n\nOptions are: \n- 'CONCURRENT' \n- 'PIPELINE'";\
	fi

go-simple-redis-store:
	go run simple_redis_store/*.go