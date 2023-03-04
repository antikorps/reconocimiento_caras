<script lang="ts">
  import Explorar from "./libs/Explorar.svelte";
  import Resultados from "./libs/Resultados.svelte";
  import Seleccionar from "./libs/Seleccionar.svelte";
  import type { analisisExploracion, respuestaProcesamiento } from "./modelos";

  let imagenMuestraRuta = "";
  let procesando = false;
  let imagenesExploradasValidas = {} as analisisExploracion[];
  let imagenesExploradasInvalidas = {} as analisisExploracion[];
  let respuestaProcesamiento = {} as respuestaProcesamiento ;

</script>

<main class="container">
  <h1>Imagen de muestra</h1>
  <article>
    <Seleccionar bind:imagenMuestraRuta bind:respuestaProcesamiento bind:procesando/>
  </article>
 

  {#if imagenMuestraRuta != ""}
  <h1>Comparar</h1>
  <article>
    <Explorar bind:imagenesExploradasValidas bind:imagenesExploradasInvalidas bind:imagenMuestraRuta bind:respuestaProcesamiento bind:procesando/>
  </article>
  {/if}

  {#if respuestaProcesamiento.resultado != null}
  <h1>Resultados</h1>
    <Resultados bind:respuestaProcesamiento/>
  {/if}

</main>

<style>
  .container {
    margin-top: 60px;
  }
</style>
