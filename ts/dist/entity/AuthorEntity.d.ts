import { PoetrydbEntityBase } from '../PoetrydbEntityBase';
import type { PoetrydbSDK } from '../PoetrydbSDK';
import type { Control } from '../types';
import type { Author, AuthorLoadMatch, AuthorListMatch } from '../PoetrydbTypes';
declare class AuthorEntity extends PoetrydbEntityBase<Author> {
    constructor(client: PoetrydbSDK, entopts: any);
    make(this: AuthorEntity): AuthorEntity;
    load(this: any, reqmatch?: AuthorLoadMatch, ctrl?: Control): Promise<AuthorEntity>;
    list(this: any, reqmatch?: AuthorListMatch, ctrl?: Control): Promise<AuthorEntity[]>;
}
export { AuthorEntity };
