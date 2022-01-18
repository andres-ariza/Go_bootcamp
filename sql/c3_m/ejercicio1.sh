# Devolver restaurante_id, nombre, barrio y tipo_cocina pero excluyendo _id para un documento (el primero).
db.restaurantes.find({}, {_id:0, grados: 0, direccion: 0})
# Devolver restaurante_id, nombre, barrio y tipo_cocina para los primeros 3 restaurantes que contengan 'Bake' en alguna parte de su nombre.
db.restaurantes.find({nombre:/Bake/}, {_id:0, grados: 0, direccion: 0})
# Contar los restaurantes de comida (tipo_cocina) china (Chinese) o tailandesa (Thai) del barrio (barrio) Bronx. Consultar or versus in.
db.restaurantes.countDocuments({ tipo_cocina: {$in: ["Chinese", "Thai"]}, barrio: "Bronx"})