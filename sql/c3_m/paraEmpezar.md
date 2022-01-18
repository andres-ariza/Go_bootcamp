Para empezar
Estas preguntas pueden responderse utilizando la interfaz gráfica de Compass.

¿Cuántas colecciones tiene la base de datos?
- tiene una llamada restaurantes

¿Cuántos documentos en cada colección? ¿Cuánto pesa cada colección?
- tiene 25K documentos, y cada uno pesa 439B para un total de 3.81MB

¿Cuántos índices en cada colección? ¿Cuánto espacio ocupan los índices de cada colección?
- hay 1 indice que ocupa 266KB

Traer un documento de ejemplo de cada colección. db.collection.find(...).pretty() nos da un formato más legible.
- usar db.restaurantes.find().pretty()

Para cada colección, listar los campos a nivel raíz (ignorar campos dentro de documentos anidados) y sus tipos de datos.
- db.restaurantes.find({}, {_id: 0, barrio: 1, "barrioType":{"$type":"$barrio"}, tipo_cocina:1, "tipo_cocinaType":{"$type":"$tipo_cocina"}, nombre:1, "nombreType":{"$type":"$nombre"}, restaurante_id:1, "restaurante_idType":{"$type":"$restaurante_id"}}) devolvera este requerimiento.