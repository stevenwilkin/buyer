# Buyer

Scheduled bitcoin buyer

## Building

	go build .

## Configuration

The [cron schedule](https://en.wikipedia.org/wiki/Cron) to use, USDT quantity to
buy each time and Binance API keys are specified via environment variables.

For convenience these values can be read from `.env` in the current directory,
see `.env.sample` for the expected format.

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
