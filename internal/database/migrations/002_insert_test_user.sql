INSERT INTO users (email, password, name) 
VALUES ('usuario@exemplo.com', 'senha123', 'Usuário Teste')
ON CONFLICT (email) DO NOTHING; 