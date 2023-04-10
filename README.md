

# CMD (Without Makefile)
- <code>go run main.go</code>
- <code>go build -o myapp</code>
- <code>./myapp</code>

# CMD (With Makefile)
- <code>make run</code>

# Test
- <code>go test ./...</code>
- <code>go test -v ./...</code>
- <code>go test ./ -run TestEqualPlayers</code> for cache
- <code>go test ./ -run TestEqualPlayers -count=1</code> for no cache
- <code>go test -v -run TestState -v -count=1 --race</code> for race condition
