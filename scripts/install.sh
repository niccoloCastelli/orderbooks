systemctl stop orderbooks
systemctl disable orderbooks

cp orderbooks /usr/bin/
cp orderbooks.service /etc/systemd/system/orderbooks.service

systemctl start orderbooks
systemctl enable orderbooks
systemctl status orderbooks

