package pkg

import (
	"fmt"
	"os"
)

func serverGenerator(dName string, pPass string) string {
	location := `location / {
			proxy_pass http://` + pPass + `;
			proxy_set_header Host $host;
			proxy_set_header X-Real-IP $remote_addr;
			proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
		}`

	server := `server {
		server_name ` + dName + `;

		` + location + `
	}`

	return server
}

func defaultService() string {
	defaultService := `server {
		listen 80 default;
		server_name _;
		return 404;
	}`

	return defaultService
}

func GenerateNginxConf(dName string, pPass string, path string) {
	fmt.Println(dName, pPass)

	location := serverGenerator(dName, pPass)

	defaultService := defaultService()

	nginxFile := fmt.Sprintf(`http {
	%s

	%s
}`, defaultService, location)

	if path != "" {
		os.Create(path)
		os.WriteFile(path+"/"+"nginx.conf", []byte(nginxFile), 0644)
	}

	os.WriteFile("nginx.conf", []byte(nginxFile), 0644)
}

func GenerateCompose(path string, imageN string, port string) {
	compose := fmt.Sprintf(`version: '3'
services:
	%s:
	  image: %s
	  volumes:
		- ./nginx.conf:/etc/nginx/nginx.conf
	  ports:
		- %s
		- 443:443
      restart: always
	`, imageN, imageN, port)

	if path != "" {
		compose := fmt.Sprintf(`version: '2'
		services:
		  %s:
			image: %s
			ports:
			  - %s
			volumes:
			  - %snginx.conf:/etc/nginx/conf.d/default.conf
			restart: always
			command: nginx -g "daemon off;"
		`, imageN, imageN, path, port)

		os.Create(path)
		os.WriteFile(path+"/"+"docker-compose.yml", []byte(compose), 0644)
		return
	}

	os.WriteFile("./docker-compose.yml", []byte(compose), 0644)
	return
}
