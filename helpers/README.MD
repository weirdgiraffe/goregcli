To be able to run insecure registry with default settings on linux with systemd do the following:

	$ sudo cp /usr/lib/systemd/system/docker.service /etc/systemd/system

And edit the line

	ExecStart=/usr/bin /docker -H fd://

To be like that

	ExecStart=/usr/bin/dockerd -H fd:// --insecure-registry=localhost:5000

And restart Docker

	$ sudo systemctl restart docker 

Now start local insecure registry:

	$ ./run-local-insecure-registry.sh
	98cd468babf41b0d08a8af4d7d539d6390129ebe3d7e8029ec9ab024a99b2d36


Now push something to our local insecure registry just to have images and tags:

	$ ./make-some-registry-entries.sh
	The push refers to a repository [localhost:5000/example]
	4fe15f8d0ae6: Pushed 
	1: digest: sha256:f11d8c9039fedc4c01435d23c75bfc955a237bb1c0df79aa613fa78e0b24e3a4 size: 528
	4fe15f8d0ae6: Layer already exists 
	2: digest: sha256:f11d8c9039fedc4c01435d23c75bfc955a237bb1c0df79aa613fa78e0b24e3a4 size: 528
	4fe15f8d0ae6: Layer already exists 
	3: digest: sha256:f11d8c9039fedc4c01435d23c75bfc955a237bb1c0df79aa613fa78e0b24e3a4 size: 528
	4fe15f8d0ae6: Layer already exists 
	latest: digest: sha256:f11d8c9039fedc4c01435d23c75bfc955a237bb1c0df79aa613fa78e0b24e3a4 size: 528
	$ curl localhost:5000/v2/_catalog
	{"repositories":["example"]}
	$ curl localhost:5000/v2/example/tags/list
	{"name":"example","tags":["1","3","2","latest"]}
