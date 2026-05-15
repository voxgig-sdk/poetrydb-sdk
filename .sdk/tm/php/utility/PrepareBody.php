<?php
declare(strict_types=1);

// Poetrydb SDK utility: prepare_body

class PoetrydbPrepareBody
{
    public static function call(PoetrydbContext $ctx): mixed
    {
        if ($ctx->op->input === 'data') {
            return ($ctx->utility->transform_request)($ctx);
        }
        return null;
    }
}
