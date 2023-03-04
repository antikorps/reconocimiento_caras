<script lang="ts">
    import type {
        analisisExploracion,
        imagenesExploradas,
        respuestaProcesamiento,
    } from "src/modelos.js";

    export let procesando: boolean;
    export let imagenesExploradasValidas: analisisExploracion[];
    export let imagenesExploradasInvalidas: analisisExploracion[];
    export let imagenMuestraRuta: string;
    export let respuestaProcesamiento: respuestaProcesamiento;
    import {
        DevolverImagenesExploradas,
        DevolverResultadoProcesamiento,
    } from "../../wailsjs/go/main/App.js";

    let errorProcesoExploracion = "";

    async function seleccionarDirectorio() {
        procesando = true
        try {
            const resultadosExploracion: imagenesExploradas =
                await DevolverImagenesExploradas();
            imagenesExploradasValidas = [];
            imagenesExploradasInvalidas = [];
            if (resultadosExploracion.validas != null) {
                for (const resultados of resultadosExploracion.validas) {
                    imagenesExploradasValidas.push(resultados);
                }
            }
            if (resultadosExploracion.invalidas != null) {
                for (const resultados of resultadosExploracion.invalidas) {
                    imagenesExploradasInvalidas.push(resultados);
                }
            }

            imagenesExploradasValidas = imagenesExploradasValidas;
            imagenesExploradasInvalidas = imagenesExploradasInvalidas;
        } catch (error) {
            imagenesExploradasValidas = [];
            imagenesExploradasInvalidas = [];
            errorProcesoExploracion = error;
        }
        procesando = false
    }

    async function procesarImagenesExploradas() {
        procesando = true
        let coleccionImagenesAnalizar = [];
        for (const imagen of imagenesExploradasValidas) {
            coleccionImagenesAnalizar.push(imagen.ruta);
        }

        respuestaProcesamiento = await DevolverResultadoProcesamiento(
            imagenMuestraRuta,
            coleccionImagenesAnalizar
        );
        procesando = false;
    }
</script>

<p>
    Selecciona el directorio o carpeta base que contiene todas las imágenes que
    quieres analizar para detectar la presencia de la imagen de muestra. La
    búsqueda será recursiva y se analizarán los subdirectorios.
</p>

<form on:submit|preventDefault={seleccionarDirectorio}>
    {#if errorProcesoExploracion != ""}
        <p>{errorProcesoExploracion}</p>
    {/if}
    <button disabled={procesando}>Selecionar directorio</button>
</form>

{#if imagenesExploradasInvalidas.length > 0}
    <h3>Archivos que se <strong>NO</strong> se analizarán:</h3>
    <ul>
        {#each imagenesExploradasInvalidas as imagen}
            <li>{imagen.ruta}: {imagen.error}</li>
        {/each}
    </ul>
{/if}

{#if imagenesExploradasValidas.length > 0}
    <h3>Imagénes que se analizarán:</h3>
    <ul>
        {#each imagenesExploradasValidas as imagen}
            <li>{imagen.ruta}</li>
        {/each}
    </ul>

    <form on:submit|preventDefault={procesarImagenesExploradas}>
        <button aria-busy={procesando}>Procesar</button>
    </form>
{/if}
