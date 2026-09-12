import { PoetrydbEntityBase } from '../PoetrydbEntityBase';
import type { PoetrydbSDK } from '../PoetrydbSDK';
import type { Control } from '../types';
import type { Random, RandomLoadMatch, RandomListMatch } from '../PoetrydbTypes';
declare class RandomEntity extends PoetrydbEntityBase<Random> {
    constructor(client: PoetrydbSDK, entopts: any);
    make(this: RandomEntity): RandomEntity;
    load(this: any, reqmatch?: RandomLoadMatch, ctrl?: Control): Promise<RandomEntity>;
    list(this: any, reqmatch?: RandomListMatch, ctrl?: Control): Promise<RandomEntity[]>;
}
export { RandomEntity };
