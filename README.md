# Shop Flowers
First way run pet-project
1. For build project need install docker.
2. Run project command <code>'docker build -t go-docker .'</code>.
3. After created the image, check it in the list with the command <code>'docker images'</code>.
4. Create and run container command <code>'docker run --rm --name shopflowers -d -p 8000:8080 go-docker'</code>.

Second way run pet-project
1. Add file docker-compose for fast start pet-project. Start command <code>'docker-compose up -d'</code>