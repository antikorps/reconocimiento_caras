<script lang="ts">
    import type { respuestaProcesamiento } from "src/modelos";
    export let respuestaProcesamiento: respuestaProcesamiento;

</script>

{#if respuestaProcesamiento.errorCritico}
    <article class="no-coincidencia">
        <p>Se ha producido un error crítico que ha impedido el análisis de las imágenes: {respuestaProcesamiento.errorCriticoMensaje}</p>
    </article>
{:else}
    {#each respuestaProcesamiento.resultado as resultado}
        {#if resultado.error}
        <article class="no-coincidencia">
            <p>Error procesando la imagen {resultado.ruta}:</p>
            <p>{resultado.errorMensaje}</p>
        </article>
        {:else if resultado.carasDetectadas == 0}
        <article class="no-coincidencia">
            <h3>{resultado.nombre}</h3>
            <p><strong>Ruta completa</strong>: {resultado.ruta}</p>
            <p>No se ha detectado ninguna coincidencia.</p>
        </article>
        {:else}
        <article class="coincidencia">
            <h3>{resultado.nombre}</h3>
            <p><strong>Ruta completa</strong>: {resultado.ruta}</p>
            {#if resultado.carasDetectadas == 1}
            <p>Se ha detectado <strong>{resultado.carasDetectadas}</strong> cara.</p>
            {:else}
            <p>Se han detectado <strong>{resultado.carasDetectadas}</strong> caras.</p>
            {/if}
            <p><img class="caras-detectadas" src={resultado.nombre} alt="cara-detectada"></p>
        </article>
        {/if}
    {/each}
{/if}

<style>
    .caras-detectadas {
        max-width: 100%;
        max-height: 300px;
        display: block;
        margin: 0 auto;
    }

    article.no-coincidencia {
        border: 6px solid tomato;
    }
    article.coincidencia {
        border: 6px solid lightgreen;
    }

</style>

