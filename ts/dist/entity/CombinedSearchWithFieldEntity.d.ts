import { PoetrydbEntityBase } from '../PoetrydbEntityBase';
import type { PoetrydbSDK } from '../PoetrydbSDK';
import type { Control } from '../types';
import type { CombinedSearchWithField, CombinedSearchWithFieldListMatch } from '../PoetrydbTypes';
declare class CombinedSearchWithFieldEntity extends PoetrydbEntityBase<CombinedSearchWithField> {
    constructor(client: PoetrydbSDK, entopts: any);
    make(this: CombinedSearchWithFieldEntity): CombinedSearchWithFieldEntity;
    list(this: any, reqmatch?: CombinedSearchWithFieldListMatch, ctrl?: Control): Promise<CombinedSearchWithFieldEntity[]>;
}
export { CombinedSearchWithFieldEntity };
