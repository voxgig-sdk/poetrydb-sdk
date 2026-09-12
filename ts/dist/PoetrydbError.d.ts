import { Context } from './Context';
declare class PoetrydbError extends Error {
    isPoetrydbError: boolean;
    sdk: string;
    code: string;
    ctx: Context;
    status: number;
    get notFound(): boolean;
    constructor(code: string, msg: string, ctx: Context);
}
export { PoetrydbError };
