export namespace explorar {
	
	export class AnalisisExploracion {
	    ruta: string;
	    error: string;
	
	    static createFrom(source: any = {}) {
	        return new AnalisisExploracion(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ruta = source["ruta"];
	        this.error = source["error"];
	    }
	}
	export class ImagenesExploradas {
	    validas: AnalisisExploracion[];
	    invalidas: AnalisisExploracion[];
	
	    static createFrom(source: any = {}) {
	        return new ImagenesExploradas(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.validas = this.convertValues(source["validas"], AnalisisExploracion);
	        this.invalidas = this.convertValues(source["invalidas"], AnalisisExploracion);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

export namespace procesar {
	
	export class ResultadoProcesamiento {
	    carasDetectadas: number;
	    error: boolean;
	    errorMensaje: string;
	    ruta: string;
	    nombre: string;
	
	    static createFrom(source: any = {}) {
	        return new ResultadoProcesamiento(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.carasDetectadas = source["carasDetectadas"];
	        this.error = source["error"];
	        this.errorMensaje = source["errorMensaje"];
	        this.ruta = source["ruta"];
	        this.nombre = source["nombre"];
	    }
	}
	export class RespuestaProcesamiento {
	    resultado: ResultadoProcesamiento[];
	    errorCritico: boolean;
	    errorCriticoMensaje: string;
	
	    static createFrom(source: any = {}) {
	        return new RespuestaProcesamiento(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.resultado = this.convertValues(source["resultado"], ResultadoProcesamiento);
	        this.errorCritico = source["errorCritico"];
	        this.errorCriticoMensaje = source["errorCriticoMensaje"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

