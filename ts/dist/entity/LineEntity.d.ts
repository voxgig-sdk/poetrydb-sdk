import { PoetrydbEntityBase } from '../PoetrydbEntityBase';
import type { PoetrydbSDK } from '../PoetrydbSDK';
import type { Control } from '../types';
import type { Line, LineLoadMatch, LineListMatch } from '../PoetrydbTypes';
declare class LineEntity extends PoetrydbEntityBase<Line> {
    constructor(client: PoetrydbSDK, entopts: any);
    make(this: LineEntity): LineEntity;
    load(this: any, reqmatch?: LineLoadMatch, ctrl?: Control): Promise<LineEntity>;
    list(this: any, reqmatch?: LineListMatch, ctrl?: Control): Promise<LineEntity[]>;
}
export { LineEntity };
