# Especialización en Back End III

## Taller de código: Desafío final

### Objetivo
A continuación se plantea un desafío integrador que nos permitirá evaluar todos los temas que hemos visto en la cursada.

### Sistema de reserva de turnos
Se desea implementar una API que permita administrar la reserva de turnos para una clínica odontológica. Esta debe cumplir con los siguientes requerimientos:

- **Administración de datos de odontólogos:** listar, agregar, modificar y eliminar odontólogos. Registrar apellido, nombre y matrícula de los mismos. Se desea el desarrollo de un CRUD para la entidad Dentista.
    - POST: agregar dentista.
    - GET: traer dentista por ID.
    - PUT: actualizar dentista.
    - PATCH: actualizar un dentista por alguno de sus campos.
    - DELETE: eliminar el dentista.

  
- **Administración de datos de los pacientes:** listar, agregar, modificar y eliminar pacientes. De cada uno se almacenan: nombre, apellido, domicilio, DNI y fecha de alta. Se desea el desarrollo de un CRUD para la entidad Paciente.
    - POST: agregar paciente.
    - GET: traer paciente por ID.
    - PUT: actualizar paciente.
    - PATCH: actualizar un paciente por alguno de sus campos.
    - DELETE: eliminar al paciente.

  

- **Registrar turno:** se tiene que poder permitir asignar a un paciente un turno con un odontólogo a una determinada fecha y hora. Al turno se le debe poder agregar una descripción. Se desea el desarrollo de un CRUD para la entidad Turno.
    - POST: agregar turno.
    - GET: traer turno por ID.
    - PUT: actualizar turno.
    - PATCH: actualizar un turno por alguno de sus campos.
    - DELETE: eliminar turno.
    - POST: agregar turno por DNI del paciente y matrícula del dentista.
    - GET: traer turno por DNI del paciente. Debe traer el detalle del turno (Fecha-Hora, descripción, Paciente y Dentista) y el DNI deberá ser recibido por QueryParams.


  
---

# Documentación de la API para el paquete 'handler' (Pacientes)

Este documento proporciona una descripción general de las funciones y endpoints disponibles en el paquete 'handler' para administrar pacientes en una clínica odontológica.

## Funciones y Endpoints

### Crear un nuevo paciente
- **Descripción:** Crea un nuevo paciente en la base de datos.
- **Cuerpo de la solicitud:** Se espera un objeto JSON con los datos del paciente.
- **Respuesta exitosa:** Devuelve el objeto del paciente creado con un código de estado 201 (Creado).
- **Errores posibles:** Si los datos de la solicitud son inválidos, se devuelve un código de estado 400 (Solicitud incorrecta). En caso de un error interno, se devuelve un código de estado 500 (Error interno del servidor).

### Obtener un paciente por ID
- **Descripción:** Crea un nuevo paciente en la base de datos.
- **Cuerpo de la solicitud:** Se espera un objeto JSON con los datos del paciente.
- **Respuesta exitosa:** Devuelve el objeto del paciente creado con un código de estado 201 (Creado).
- **Errores posibles:** Si los datos de la solicitud son inválidos, se devuelve un código de estado 400 (Solicitud incorrecta). En caso de un error interno, se devuelve un código de estado 500 (Error interno del servidor).

### Obtener un paciente por ID
- **Descripción:** Obtiene un paciente específico según su ID.
- **Parámetros de la URL:** Se espera el ID del paciente.
- **Respuesta exitosa:** Devuelve los detalles del paciente solicitado con un código de estado 200 (Éxito).
- **Errores posibles:** Si el ID del paciente es inválido o el paciente no se encuentra, se devuelve un código de estado 400 (Solicitud incorrecta) o 404 (No encontrado), respectivamente.

### Actualizar un paciente
- **Descripción:** Actualiza un paciente existente según su ID.
- **Parámetros de la URL:** Se espera el ID del paciente a actualizar.
- **Cuerpo de la solicitud:** Se espera un objeto JSON con los datos actualizados del paciente.
- **Respuesta exitosa:** Devuelve el objeto del paciente actualizado con un código de estado 200 (Éxito).
- **Errores posibles:** Si el ID del paciente es inválido, los datos de la solicitud son incorrectos o hay un error interno, se devuelven códigos de estado 400 (Solicitud incorrecta) o 500 (Error interno del servidor), respectivamente.

### Actualizar parcialmente un paciente
- **Descripción:** Actualiza parcialmente un paciente existente según su ID.
- **Parámetros de la URL:** Se espera el ID del paciente a actualizar.
- **Cuerpo de la solicitud:** Se espera un objeto JSON con los campos a actualizar del paciente.
- **Respuesta exitosa:** Devuelve el objeto del paciente actualizado con un código de estado 200 (Éxito).
- **Errores posibles:** Si el ID del paciente es inválido, los datos de la solicitud son incorrectos o hay un error interno, se devuelven códigos de estado 400 (Solicitud incorrecta) o 500 (Error interno del servidor), respectivamente.

### Eliminar un paciente
- **Descripción:** Elimina un paciente existente según su ID.
- **Parámetros de la URL:** Se espera el ID del paciente a eliminar.
- **Respuesta exitosa:** Devuelve un mensaje de éxito con un código de estado 200 (Éxito).
- **Errores posibles:** Si el ID del paciente es inválido o el paciente no se encuentra, se devuelven códigos de estado 400 (Solicitud incorrecta) o 404 (No encontrado), respectivamente.

Esta documentación proporciona información sobre las funciones y endpoints disponibles en el paquete 'handler' para administrar pacientes en una clínica odontológica. Asegúrate de utilizar estos endpoints según sea necesario para interactuar con la API.


----

# Documentación de la API para el paquete 'handler' (Dentistas)

Este documento proporciona una descripción general de las funciones y endpoints disponibles en el paquete 'handler' para administrar dentistas en una clínica odontológica.

## Funciones y Endpoints

### Crear un nuevo dentista
- **Descripción:** Crea un nuevo dentista en la base de datos.
- **Cuerpo de la solicitud:** Se espera un objeto JSON con los datos del dentista.
- **Respuesta exitosa:** Devuelve el objeto del dentista creado con un código de estado 201 (Creado).
- **Errores posibles:** Si los datos de la solicitud son inválidos, se devuelve un código de estado 400 (Solicitud incorrecta). En caso de un error interno, se devuelve un código de estado 500 (Error interno del servidor).

### Obtener un dentista por ID
- **Descripción:** Obtiene un dentista específico según su ID.
- **Parámetros de la URL:** Se espera el ID del dentista.
- **Respuesta exitosa:** Devuelve los detalles del dentista solicitado con un código de estado 200 (Éxito).
- **Errores posibles:** Si el ID del dentista es inválido o el dentista no se encuentra, se devuelve un código de estado 400 (Solicitud incorrecta) o 404 (No encontrado), respectivamente.

### Actualizar un dentista
- **Descripción:** Actualiza un dentista existente según su ID.
- **Parámetros de la URL:** Se espera el ID del dentista a actualizar.
- **Cuerpo de la solicitud:** Se espera un objeto JSON con los datos actualizados del dentista.
- **Respuesta exitosa:** Devuelve el objeto del dentista actualizado con un código de estado 200 (Éxito).
- **Errores posibles:** Si el ID del dentista es inválido, los datos de la solicitud son incorrectos o hay un error interno, se devuelven códigos de estado 400 (Solicitud incorrecta) o 500 (Error interno del servidor), respectivamente.

### Actualizar parcialmente un dentista
- **Descripción:** Actualiza parcialmente un dentista existente según su ID.
- **Parámetros de la URL:** Se espera el ID del dentista a actualizar.
- **Cuerpo de la solicitud:** Se espera un objeto JSON con los campos a actualizar del dentista.
- **Respuesta exitosa:** Devuelve el objeto del dentista actualizado con un código de estado 200 (Éxito).
- **Errores posibles:** Si el ID del dentista es inválido, los datos de la solicitud son incorrectos o hay un error interno, se devuelven códigos de estado 400 (Solicitud incorrecta) o 500 (Error interno del servidor), respectivamente.

### Eliminar un dentista
- **Descripción:** Elimina un dentista existente según su ID.
- **Parámetros de la URL:** Se espera el ID del dentista a eliminar.
- **Respuesta exitosa:** Devuelve un mensaje de éxito con un código de estado 200 (Éxito).
- **Errores posibles:** Si el ID del dentista es inválido o el dentista no se encuentra, se devuelven códigos de estado 400 (Solicitud incorrecta) o 404 (No encontrado), respectivamente.

Esta documentación proporciona información sobre las funciones y endpoints disponibles en el paquete 'handler' para administrar dentistas en una clínica odontológica. Asegúrate de utilizar estos endpoints según sea necesario para interactuar con la API.


---
# Documentación de la API para el paquete 'handler' (Appointments)

Este documento proporciona una descripción general de las funciones y endpoints disponibles en el paquete 'handler' para administrar citas médicas en una clínica odontológica.

## Funciones y Endpoints

### Crear una nueva cita médica


### Crear una nueva cita médica
- **Descripción:** Crea una nueva cita médica en la base de datos.
- **Cuerpo de la solicitud:** Se espera un objeto JSON con los datos de la cita médica.
- **Respuesta exitosa:** Devuelve el objeto de cita médica creado con un código de estado 201 (Creado).
- **Errores posibles:** Si los datos de la solicitud son inválidos, se devuelve un código de estado 400 (Solicitud incorrecta). En caso de un error interno, se devuelve un código de estado 500 (Error interno del servidor).

### Obtener una cita médica por ID
- **Descripción:** Obtiene una cita médica específica según su ID.
- **Parámetros de la URL:** Se espera el ID de la cita médica.
- **Respuesta exitosa:** Devuelve los detalles de la cita médica solicitada con un código de estado 200 (Éxito).
- **Errores posibles:** Si el ID de la cita es inválido o la cita no se encuentra, se devuelve un código de estado 400 (Solicitud incorrecta) o 404 (No encontrado), respectivamente.

### Actualizar una cita médica
- **Descripción:** Actualiza una cita médica existente según su ID.
- **Parámetros de la URL:** Se espera el ID de la cita médica a actualizar.
- **Cuerpo de la solicitud:** Se espera un objeto JSON con los datos actualizados de la cita médica.
- **Respuesta exitosa:** Devuelve el objeto de cita médica actualizado con un código de estado 200 (Éxito).
- **Errores posibles:** Si el ID de la cita es inválido, los datos de la solicitud son incorrectos o hay un error interno, se devuelven códigos de estado 400 (Solicitud incorrecta) o 500 (Error interno del servidor), respectivamente.

### Actualizar parcialmente una cita médica
- **Descripción:** Actualiza parcialmente una cita médica existente según su ID.
- **Parámetros de la URL:** Se espera el ID de la cita médica a actualizar.
- **Cuerpo de la solicitud:** Se espera un objeto JSON con los campos a actualizar de la cita médica.
- **Respuesta exitosa:** Devuelve el objeto de cita médica actualizado con un código de estado 200 (Éxito).
- **Errores posibles:** Si el ID de la cita es inválido, los datos de la solicitud son incorrectos o hay un error interno, se devuelven códigos de estado 400 (Solicitud incorrecta) o 500 (Error interno del servidor), respectivamente.

### Eliminar una cita médica
- **Descripción:** Elimina una cita médica existente según su ID.
- **Parámetros de la URL:** Se espera el ID de la cita médica a eliminar.
- **Respuesta exitosa:** Devuelve un mensaje de éxito con un código de estado 200 (Éxito).
- **Errores posibles:** Si el ID de la cita es inválido o la cita no se encuentra, se devuelve un código de estado 400 (Solicitud incorrecta) o 404 (No encontrado), respectivamente.

### Obtener citas médicas por paciente
- **Descripción:** Obtiene todas las citas médicas asociadas a un paciente según su ID de credencial.
- **Parámetros de la URL:** Se espera el ID de la credencial del paciente.
- **Respuesta exitosa:** Devuelve una lista de citas médicas relacionadas con el paciente con un código de estado 200 (Éxito).
- **Errores posibles:** Si el ID de la credencial es inválido o no se encuentra ninguna cita asociada al paciente, se devuelve un código de estado 400 (Solicitud incorrecta) o 404 (No encontrado), respectivamente.



