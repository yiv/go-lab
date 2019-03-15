package mysql

const TableConsumersTemplate = `
  id int(10) unsigned NOT NULL AUTO_INCREMENT,
  consumerid char(36) NOT NULL DEFAULT '',
  position int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (id),
  UNIQUE KEY consumerid (consumerid)
`

