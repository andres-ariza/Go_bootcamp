# Traer 3 restaurantes que hayan recibido al menos una calificación de grado 'A' con puntaje mayor a 20. 
# Una misma calificación debe cumplir con ambas condiciones simultáneamente; investigar el operador elemMatch.
db.restaurantes.find({ grados: { $elemMatch:  { grado: "A", puntaje: {$gt: 20}}}}).limit(3)
# ¿A cuántos documentos les faltan las coordenadas geográficas? En otras palabras, revisar si el tamaño de direccion.coord es 0 y contar.
db.restaurantes.countDocuments({"direccion.coord":{$size:0}})
# Devolver nombre, barrio, tipo_cocina y grados para los primeros 3 restaurantes; de cada documento solo la última calificación.
# Ver el operador slice.
db.restaurantes.find({}, {nombre:1, barrio:1, tipo_cocina:1, grados: {$slice: -1}}).limit(3)