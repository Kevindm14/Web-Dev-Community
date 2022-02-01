# PetStore Loneliness 2000

PetStore Loneliness 2000 es un negocio que se dedica a la venta y compra de mascotas. Al ser recientemente fundados, una de sus necesidades primordiales es la de un sistema de información que gestione los datos de las mascotas y los clientes. Los expertos en TI establecen que la aplicación constará de dos capas: la lógica del servidor y la base de datos.

Como primer paso, el sistema necesita gestionar los datos de las mascotas y asociar los siguientes campos a cada una:
- Animal (perro, gato, etc.)
- Raza
- Precio
- Edad

Como segundo paso, y de la misma forma que las mascotas, se necesita gestionar los datos de los clientes y asociar los siguientes campos a cada uno:

- Nombre
- Apellido
- Email
- Telefono
- Direccion
- Género
- Edad

Por gestionar la información de estas entidades se quiere decir:
- Que se puedan guardar de forma permanente en la BDD.
- Que muestre todas las mascotas/clientes y su información.
- Que se puedan filtrar las mascotas de acuerdo a los siguientes criterios:
    - Animal. Ejemplo: mostrar las mascotas que sean "perros" o mostrar las que sean "gatos".
    - Precio (Hasta X monto inclusive). Ejemplo :mostrar las mascotas que estén por debajo de $700.000. (Podría validarse si la búsqueda contiene sólo caracteres numéricos)
    - Edad (meses o años): mostrar las mascotas que tengan X meses o Y años (es responsabilidad del usuario especificar si es mes o año)
- Que se puedan filtrar los clientes de acuerdo a los siguientes criterios:
    - Nombre y/o Apellido. Ejemplo: si un cliente se llama Pepo Aguilar, podría mostrarse como parte de los resultados si ingreso la búsqueda como "Pepo" o "aguilar" o "pepo aguilar".
    - Email. (Podría validarse si la búsqueda contiene el formato adecuado para emails).
    - Qué mascotas tiene un cliente? Debe mostrar la información del cliente junto con el animal y la raza.
- Que se puedan remover o eliminar entidades.
- Que se puedan actualizar entidades.

Tener en cuenta que un cliente está registrado siempre y cuando esté asociado con al menos una mascota (al momento de la compra se verifica si existe el cliente).