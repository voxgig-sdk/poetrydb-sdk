<?php
declare(strict_types=1);

// Poetrydb SDK utility: result_body

class PoetrydbResultBody
{
    public static function call(PoetrydbContext $ctx): ?PoetrydbResult
    {
        $response = $ctx->response;
        $result = $ctx->result;
        if ($result && $response && $response->json_func && $response->body) {
            $result->body = ($response->json_func)();
        }
        return $result;
    }
}
