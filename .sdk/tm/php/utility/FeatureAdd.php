<?php
declare(strict_types=1);

// Poetrydb SDK utility: feature_add

class PoetrydbFeatureAdd
{
    public static function call(PoetrydbContext $ctx, mixed $f): void
    {
        $ctx->client->features[] = $f;
    }
}
