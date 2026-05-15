# Poetrydb SDK utility: make_context

from core.context import PoetrydbContext


def make_context_util(ctxmap, basectx):
    return PoetrydbContext(ctxmap, basectx)
