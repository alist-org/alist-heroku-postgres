mkdir -p /opt/alist/data/
# DATABASE_URL = postgres://user3123:passkja83kd8@ec2-117-21-174-214.compute-1.amazonaws.com:6212/db982398
/main
cat >/opt/alist/data/config.json <<EOF
{
  "address": "0.0.0.0",
  "port": $PORT,
  "assets": "$ASSETS",
  "database": {
    "type": "postgres",
    "table_prefix": "x_",
    "ssl_mode": "require"
  },
  "scheme": {
    "https": false,
    "cert_file": "",
    "key_file": ""
  },
  "cache": {
    "expiration": $EXPIRATION,
    "cleanup_interval": $CLEANUP_INTERVAL
  },
  "temp_dir": "data/temp"
}
EOF

cd /opt/alist
./alist -conf data/config.json -docker