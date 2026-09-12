import { PoetrydbEntityBase } from '../PoetrydbEntityBase';
import type { PoetrydbSDK } from '../PoetrydbSDK';
import type { Control } from '../types';
import type { Title, TitleLoadMatch, TitleListMatch } from '../PoetrydbTypes';
declare class TitleEntity extends PoetrydbEntityBase<Title> {
    constructor(client: PoetrydbSDK, entopts: any);
    make(this: TitleEntity): TitleEntity;
    load(this: any, reqmatch?: TitleLoadMatch, ctrl?: Control): Promise<TitleEntity>;
    list(this: any, reqmatch?: TitleListMatch, ctrl?: Control): Promise<TitleEntity[]>;
}
export { TitleEntity };
