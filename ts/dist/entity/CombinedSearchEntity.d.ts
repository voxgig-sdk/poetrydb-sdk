import { PoetrydbEntityBase } from '../PoetrydbEntityBase';
import type { PoetrydbSDK } from '../PoetrydbSDK';
import type { Control } from '../types';
import type { CombinedSearch, CombinedSearchListMatch } from '../PoetrydbTypes';
declare class CombinedSearchEntity extends PoetrydbEntityBase<CombinedSearch> {
    constructor(client: PoetrydbSDK, entopts: any);
    make(this: CombinedSearchEntity): CombinedSearchEntity;
    list(this: any, reqmatch?: CombinedSearchListMatch, ctrl?: Control): Promise<CombinedSearchEntity[]>;
}
export { CombinedSearchEntity };
