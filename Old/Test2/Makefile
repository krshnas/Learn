#### Mock Server #####

## TODO: Remove this when Registration Service interaction is ready
mockserver:
	helm upgrade --create-namespace --install mockserver mockserver/mockserver \
		--version 5.14.0 \
		--namespace mockserver

mockserver-:
	helm uninstall --ignore-not-found -n mockserver mockserver
	kubectl delete namespace --ignore-not-found mockserver
