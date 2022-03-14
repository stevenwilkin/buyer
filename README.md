# Buyer

Scheduled bitcoin buyer

## Building

	go build .

## Running

	./buyer

## Systemd service

Copy the service unit file to the configuration directory:

	cp buyer.service /etc/systemd/system

Enable and start the service:

	systemctl enable buyer
	systemctl start buyer

## Rsyslog forwarding

	cat << EOF > /etc/rsyslog.d/20-buyer.conf
	if $programname == 'buyer' then @HOST.papertrailapp.com:PORT
	& ~
	EOF
	systemctl restart rsyslog
