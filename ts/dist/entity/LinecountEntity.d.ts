import { PoetrydbEntityBase } from '../PoetrydbEntityBase';
import type { PoetrydbSDK } from '../PoetrydbSDK';
import type { Control } from '../types';
import type { Linecount, LinecountLoadMatch, LinecountListMatch } from '../PoetrydbTypes';
declare class LinecountEntity extends PoetrydbEntityBase<Linecount> {
    constructor(client: PoetrydbSDK, entopts: any);
    make(this: LinecountEntity): LinecountEntity;
    load(this: any, reqmatch?: LinecountLoadMatch, ctrl?: Control): Promise<LinecountEntity>;
    list(this: any, reqmatch?: LinecountListMatch, ctrl?: Control): Promise<LinecountEntity[]>;
}
export { LinecountEntity };
