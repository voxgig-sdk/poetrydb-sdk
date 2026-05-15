<?php
declare(strict_types=1);

// Poetrydb SDK utility: make_context

require_once __DIR__ . '/../core/Context.php';

class PoetrydbMakeContext
{
    public static function call(array $ctxmap, ?PoetrydbContext $basectx): PoetrydbContext
    {
        return new PoetrydbContext($ctxmap, $basectx);
    }
}
