import { AuthorEntity } from './entity/AuthorEntity';
import { AuthorabEntity } from './entity/AuthorabEntity';
import { CombinedSearchEntity } from './entity/CombinedSearchEntity';
import { CombinedSearchWithFieldEntity } from './entity/CombinedSearchWithFieldEntity';
import { LineEntity } from './entity/LineEntity';
import { LinecountEntity } from './entity/LinecountEntity';
import { PoemcountEntity } from './entity/PoemcountEntity';
import { RandomEntity } from './entity/RandomEntity';
import { TitleEntity } from './entity/TitleEntity';
import { TitleabEntity } from './entity/TitleabEntity';
export type * from './PoetrydbTypes';
import { inspect } from 'node:util';
import type { Context, Feature } from './types';
import { config } from './Config';
import { PoetrydbEntityBase } from './PoetrydbEntityBase';
import { Utility } from './utility/Utility';
import { BaseFeature } from './feature/base/BaseFeature';
declare const stdutil: Utility;
declare class PoetrydbSDK {
    _mode: string;
    _options: any;
    _utility: Utility;
    _features: Feature[];
    _rootctx: Context;
    constructor(options?: any);
    options(): any;
    utility(): any;
    prepare(fetchargs?: any): Promise<any>;
    direct(fetchargs?: any): Promise<Error | {
        ok: boolean;
        status: number;
        headers: any;
        data: any;
        err?: undefined;
    } | {
        ok: boolean;
        err: any;
        status?: undefined;
        headers?: undefined;
        data?: undefined;
    }>;
    _rawRequest(fetchargs?: any): Promise<Error | {
        ok: boolean;
        status: number;
        headers: any;
        data: any;
        err?: undefined;
    } | {
        ok: boolean;
        err: any;
        status?: undefined;
        headers?: undefined;
        data?: undefined;
    }>;
    graphql(query: string, variables?: any, ctrl?: any): Promise<any>;
    Author(entopts?: Record<string, any>): AuthorEntity;
    Authorab(entopts?: Record<string, any>): AuthorabEntity;
    CombinedSearch(entopts?: Record<string, any>): CombinedSearchEntity;
    CombinedSearchWithField(entopts?: Record<string, any>): CombinedSearchWithFieldEntity;
    Line(entopts?: Record<string, any>): LineEntity;
    Linecount(entopts?: Record<string, any>): LinecountEntity;
    Poemcount(entopts?: Record<string, any>): PoemcountEntity;
    Random(entopts?: Record<string, any>): RandomEntity;
    Title(entopts?: Record<string, any>): TitleEntity;
    Titleab(entopts?: Record<string, any>): TitleabEntity;
    static test(testoptsarg?: any, sdkoptsarg?: any): PoetrydbSDK;
    tester(testopts?: any, sdkopts?: any): PoetrydbSDK;
    toJSON(): {
        name: string;
    };
    toString(): string;
    [inspect.custom](): string;
}
declare const SDK: typeof PoetrydbSDK;
export { stdutil, config, BaseFeature, PoetrydbEntityBase, PoetrydbSDK, SDK, };
