CREATE TABLE IF NOT EXISTS relatorios (
  cpf VARCHAR(14) PRIMARY KEY,
  private BOOLEAN,
  incompleto BOOLEAN,
  data_ultima_compra DATE,
  ticket_medio DECIMAL(10,2),
  ticket_ultima_compra DECIMAL(10,2),
  loja_mais_frequente VARCHAR(14),
  loja_ultima_compra VARCHAR(14)
);