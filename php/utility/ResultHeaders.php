<?php
declare(strict_types=1);

// Poetrydb SDK utility: result_headers

class PoetrydbResultHeaders
{
    public static function call(PoetrydbContext $ctx): ?PoetrydbResult
    {
        $response = $ctx->response;
        $result = $ctx->result;
        if ($result) {
            if ($response && is_array($response->headers)) {
                $result->headers = $response->headers;
            } else {
                $result->headers = [];
            }
        }
        return $result;
    }
}
