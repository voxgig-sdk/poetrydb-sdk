import { PoetrydbEntityBase } from '../PoetrydbEntityBase';
import type { PoetrydbSDK } from '../PoetrydbSDK';
import type { Control } from '../types';
import type { Titleab, TitleabListMatch } from '../PoetrydbTypes';
declare class TitleabEntity extends PoetrydbEntityBase<Titleab> {
    constructor(client: PoetrydbSDK, entopts: any);
    make(this: TitleabEntity): TitleabEntity;
    list(this: any, reqmatch?: TitleabListMatch, ctrl?: Control): Promise<TitleabEntity[]>;
}
export { TitleabEntity };
