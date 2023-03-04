<script lang="ts">
    export let imagenMuestraRuta;
    import type { respuestaProcesamiento } from "src/modelos.js";
    import { DevolverImagenReconocimiento } from "../../wailsjs/go/main/App.js";
    export let respuestaProcesamiento: respuestaProcesamiento;
    export let procesando: boolean;

    let mensajeErrorImagenMuestra = "";
    async function seleccionarImagenMuestra() {
        procesando = true
        try {
            imagenMuestraRuta = await DevolverImagenReconocimiento();
        } catch (error) {
            imagenMuestraRuta = "";
            mensajeErrorImagenMuestra = error;
        }
        respuestaProcesamiento = {} as respuestaProcesamiento;
        procesando = false
    }
</script>

<form on:submit|preventDefault={seleccionarImagenMuestra}>
    {#if imagenMuestraRuta != ""}
        <img
            src={imagenMuestraRuta}
            alt="Imagen de muestra"
            id="imagen-muestra"
        />
    {:else}
        <svg
            xmlns="http://www.w3.org/2000/svg"
            id="imagen-base"
            viewBox="0 0 180.119 139.794"
            ><g
                transform="translate(-13.59 -66.639)"
                paint-order="fill markers stroke"
                ><path
                    fill="#d0d0d0"
                    d="M13.591 66.639H193.71v139.794H13.591z"
                /><path
                    d="m118.507 133.514-34.249 34.249-15.968-15.968-41.938 41.937H178.726z"
                    opacity=".675"
                    fill="#fff"
                /><circle
                    cx="58.217"
                    cy="108.555"
                    r="11.773"
                    opacity=".675"
                    fill="#fff"
                /><path
                    fill="none"
                    d="M26.111 77.634h152.614v116.099H26.111z"
                /></g
            ></svg
        >
        <p id="mensajeError">{mensajeErrorImagenMuestra}</p>
    {/if}

    <p>
        Selecciona la fotografía de la persona que se quiere buscar en la
        colección de imágenes.
    </p>
    <button type="submit" disabled={procesando}>Seleccionar</button>
</form>

<style>
    #imagen-base,
    #imagen-muestra {
        max-width: 100%;
        max-height: 150px;
        display: block;
        margin: 20px auto 20px auto;
    }
    form,
    #mensajeError {
        margin: 20px 0 20px 0;
    }
</style>
