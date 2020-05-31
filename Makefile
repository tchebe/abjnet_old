

dockerbuild_user_service:
	cd user_service && docker build -t abjnet_user_service . 

dockerbuild_user_service_proxy:
	cd user_service && docker build --build-arg HTTP_PROXY=http://172.17.0.1:3128/ --build-arg HTTPS_PROXY=http://172.17.0.1:3128/ -t abjnet_user_service . 

dockerbuild_email_service:
	cd .. && cd email_service && docker build -t abjnet_email_service . 

dockerbuild_email_service_proxy:
	cd .. && cd email_service && docker build --build-arg HTTP_PROXY=http://172.17.0.1:3128/  --build-arg HTTPS_PROXY=http://172.17.0.1:3128/ -t abjnet_email_service . 

dockerbuild_product_service:
	cd .. && cd product_service && docker build -t abjnet_product_service . 

dockerbuild_product_service_proxy:
	cd .. && cd product_service && docker build --build-arg HTTP_PROXY=http://172.17.0.1:3128/  --build-arg HTTPS_PROXY=http://172.17.0.1:3128/ -t abjnet_product_service . 

dockerbuild_souscription_service:
	cd .. && cd souscription_service && docker build -t abjnet_souscription_service . 

dockerbuild_souscription_service_proxy:
	cd .. && cd souscription_service && docker build --build-arg HTTP_PROXY=http://172.17.0.1:3128/  --build-arg HTTPS_PROXY=http://172.17.0.1:3128/ -t abjnet_souscription_service . 

dockerbuild_taskrunner_service:
	cd .. && cd taskrunner_service && docker build -t abjnet_taskrunner_service . 

dockerbuild_taskrunner_service_proxy:
	cd .. && cd taskrunner_service && docker build --build-arg HTTP_PROXY=http://172.17.0.1:3128/  --build-arg HTTPS_PROXY=http://172.17.0.1:3128/ -t abjnet_taskrunner_service . 

dockerbuild_prestation_service:
	cd .. && cd prestation_service && docker build -t abjnet_prestation_service . 

dockerbuild_prestation_service_proxy:
	cd .. && cd prestation_service && docker build --build-arg HTTP_PROXY=http://172.17.0.1:3128/  --build-arg HTTPS_PROXY=http://172.17.0.1:3128/ -t abjnet_prestation_service . 

dockerbuild_payment_service:
	cd .. && cd payment_service && docker build -t abjnet_payment_service . 

dockerbuild_payment_service_proxy:
	cd .. && cd payment_service && docker build --build-arg HTTP_PROXY=http://172.17.0.1:3128/  --build-arg HTTPS_PROXY=http://172.17.0.1:3128/ -t abjnet_payment_service . 

dockerbuild_api_user:
	cd .. && cd restapi && docker build -t abjnet_api_rest . 

dockerbuild_api_user_proxy:
	cd .. && cd restapi && docker build --build-arg HTTP_PROXY=http://172.17.0.1:3128/  --build-arg HTTPS_PROXY=http://172.17.0.1:3128/ -t abjnet_api_rest . 

dockerbuild_api_gateway:
	cd .. && cd api_gateway && docker build -t abjnet_api_gateway . 

dockerbuild_api_gateway_proxy:
	cd .. && cd api_gateway && docker build --build-arg HTTP_PROXY=http://172.17.0.1:3128/  --build-arg HTTPS_PROXY=http://172.17.0.1:3128/ -t abjnet_api_gateway . 

dockerbuild_all:
	dockerbuild_user_service dockerbuild_email_service dockerbuild_product_service dockerbuild_souscription_service dockerbuild_taskrunner_service dockerbuild_prestation_service dockerbuild_payment_service dockerbuild_api_gateway