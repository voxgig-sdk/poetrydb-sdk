import { PoetrydbEntityBase } from '../PoetrydbEntityBase';
import type { PoetrydbSDK } from '../PoetrydbSDK';
import type { Control } from '../types';
import type { Poemcount, PoemcountLoadMatch } from '../PoetrydbTypes';
declare class PoemcountEntity extends PoetrydbEntityBase<Poemcount> {
    constructor(client: PoetrydbSDK, entopts: any);
    make(this: PoemcountEntity): PoemcountEntity;
    load(this: any, reqmatch?: PoemcountLoadMatch, ctrl?: Control): Promise<PoemcountEntity>;
}
export { PoemcountEntity };
