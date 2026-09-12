import { PoetrydbEntityBase } from '../PoetrydbEntityBase';
import type { PoetrydbSDK } from '../PoetrydbSDK';
import type { Control } from '../types';
import type { Authorab, AuthorabListMatch } from '../PoetrydbTypes';
declare class AuthorabEntity extends PoetrydbEntityBase<Authorab> {
    constructor(client: PoetrydbSDK, entopts: any);
    make(this: AuthorabEntity): AuthorabEntity;
    list(this: any, reqmatch?: AuthorabListMatch, ctrl?: Control): Promise<AuthorabEntity[]>;
}
export { AuthorabEntity };
