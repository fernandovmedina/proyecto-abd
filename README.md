Cree una base de datos llamada 'AUTH'. (Solo para mañana, despues la integramos a la base de datos)
Cree una tabla llamada USERS con los siguientes campos:
  - ID INT AUTO_INCREMENT NOT NULL,
  - NAME TEXT NULL,
  - LASTNAME TEXT NULL,
  - EMAIL TEXT NULL,
  - PASS TEXT NULL

Cambia los valores en el archivo .env, principalmente el db_user y el db_pass, checa que el host y port sean correctos ya que usas windows.

Por ultimo ejecuta el binario llamado "bin", creo que en windows se hace asi "\.bin" o podria ser asi "./bin"

Ya deberia de funcionar aqui, abre el navegador en http://127.0.0.1:8080/register o http://localhost:8080/register
