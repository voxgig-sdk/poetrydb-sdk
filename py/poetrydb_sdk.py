# Poetrydb SDK

from utility.voxgig_struct import voxgig_struct as vs
from core.utility_type import PoetrydbUtility
from core.spec import PoetrydbSpec
from core import helpers

# Load utility registration (populates Utility._registrar)
from utility import register

# Load features
from feature.base_feature import PoetrydbBaseFeature
from features import _make_feature


class PoetrydbSDK:

    def __init__(self, options=None):
        self.mode = "live"
        self.features = []
        self.options = None

        utility = PoetrydbUtility()
        self._utility = utility

        from config import make_config
        config = make_config()

        self._rootctx = utility.make_context({
            "client": self,
            "utility": utility,
            "config": config,
            "options": options if options is not None else {},
            "shared": {},
        }, None)

        self.options = utility.make_options(self._rootctx)

        if vs.getpath(self.options, "feature.test.active") is True:
            self.mode = "test"

        self._rootctx.options = self.options

        # Add features from config.
        feature_opts = helpers.to_map(vs.getprop(self.options, "feature"))
        if feature_opts is not None:
            feature_items = vs.items(feature_opts)
            if feature_items is not None:
                for item in feature_items:
                    fname = item[0]
                    fopts = helpers.to_map(item[1])
                    if fopts is not None and fopts.get("active") is True:
                        utility.feature_add(self._rootctx, _make_feature(fname))

        # Add extension features.
        extend = vs.getprop(self.options, "extend")
        if isinstance(extend, list):
            for f in extend:
                if isinstance(f, dict) or (hasattr(f, "get_name") and callable(f.get_name)):
                    utility.feature_add(self._rootctx, f)

        # Initialize features.
        for f in self.features:
            utility.feature_init(self._rootctx, f)

        utility.feature_hook(self._rootctx, "PostConstruct")

        # #BuildFeatures

    def options_map(self):
        out = vs.clone(self.options)
        if isinstance(out, dict):
            return out
        return {}

    def get_utility(self):
        return PoetrydbUtility.copy(self._utility)

    def get_root_ctx(self):
        return self._rootctx

    def prepare(self, fetchargs=None):
        utility = self._utility

        if fetchargs is None:
            fetchargs = {}

        ctrl = helpers.to_map(vs.getprop(fetchargs, "ctrl"))
        if ctrl is None:
            ctrl = {}

        ctx = utility.make_context({
            "opname": "prepare",
            "ctrl": ctrl,
        }, self._rootctx)

        options = self.options

        path = vs.getprop(fetchargs, "path") or ""
        if not isinstance(path, str):
            path = ""

        method = vs.getprop(fetchargs, "method") or "GET"
        if not isinstance(method, str):
            method = "GET"

        params = helpers.to_map(vs.getprop(fetchargs, "params"))
        if params is None:
            params = {}
        query = helpers.to_map(vs.getprop(fetchargs, "query"))
        if query is None:
            query = {}

        headers = utility.prepare_headers(ctx)

        base = vs.getprop(options, "base") or ""
        if not isinstance(base, str):
            base = ""
        prefix = vs.getprop(options, "prefix") or ""
        if not isinstance(prefix, str):
            prefix = ""
        suffix = vs.getprop(options, "suffix") or ""
        if not isinstance(suffix, str):
            suffix = ""

        ctx.spec = PoetrydbSpec({
            "base": base,
            "prefix": prefix,
            "suffix": suffix,
            "path": path,
            "method": method,
            "params": params,
            "query": query,
            "headers": headers,
            "body": vs.getprop(fetchargs, "body"),
            "step": "start",
        })

        # Merge user-provided headers.
        uh = vs.getprop(fetchargs, "headers")
        if isinstance(uh, dict):
            for k, v in uh.items():
                ctx.spec.headers[k] = v

        _, err = utility.prepare_auth(ctx)
        if err is not None:
            raise err

        fetchdef, err = utility.make_fetch_def(ctx)
        if err is not None:
            raise err

        return fetchdef

    def direct(self, fetchargs=None):
        utility = self._utility

        try:
            fetchdef = self.prepare(fetchargs)
        except Exception as err:
            # direct() is the raw-HTTP escape hatch: it never raises, it
            # returns a result object callers branch on via result["ok"].
            return {"ok": False, "err": err}

        if fetchargs is None:
            fetchargs = {}
        ctrl = helpers.to_map(vs.getprop(fetchargs, "ctrl"))
        if ctrl is None:
            ctrl = {}

        ctx = utility.make_context({
            "opname": "direct",
            "ctrl": ctrl,
        }, self._rootctx)

        url = fetchdef.get("url", "")
        fetched, fetch_err = utility.fetcher(ctx, url, fetchdef)

        if fetch_err is not None:
            return {"ok": False, "err": fetch_err}

        if fetched is None:
            return {
                "ok": False,
                "err": ctx.make_error("direct_no_response", "response: undefined"),
            }

        if isinstance(fetched, dict):
            status = helpers.to_int(vs.getprop(fetched, "status"))
            headers = vs.getprop(fetched, "headers") or {}

            # No-body responses (204, 304) and explicit zero content-length
            # must skip JSON parsing — calling json() on an empty body raises.
            content_length = None
            if isinstance(headers, dict):
                content_length = headers.get("content-length")
            no_body = status in (204, 304) or str(content_length) == "0"

            json_data = None
            if not no_body:
                jf = vs.getprop(fetched, "json")
                if callable(jf):
                    try:
                        json_data = jf()
                    except Exception:
                        # Non-JSON body (e.g. text/plain, text/html). Surface
                        # status + headers but leave data as None.
                        json_data = None

            return {
                "ok": status >= 200 and status < 300,
                "status": status,
                "headers": headers,
                "data": json_data,
            }

        return {
            "ok": False,
            "err": ctx.make_error("direct_invalid", "invalid response type"),
        }


    @property
    def author(self):
        """Idiomatic facade: client.author.list() / client.author.load({"id": ...})."""
        from entity.author_entity import AuthorEntity
        cached = getattr(self, "_author", None)
        if cached is None:
            cached = AuthorEntity(self, None)
            self._author = cached
        return cached

    def Author(self, data=None):
        # Deprecated: use client.author instead.
        from entity.author_entity import AuthorEntity
        return AuthorEntity(self, data)


    @property
    def authorab(self):
        """Idiomatic facade: client.authorab.list() / client.authorab.load({"id": ...})."""
        from entity.authorab_entity import AuthorabEntity
        cached = getattr(self, "_authorab", None)
        if cached is None:
            cached = AuthorabEntity(self, None)
            self._authorab = cached
        return cached

    def Authorab(self, data=None):
        # Deprecated: use client.authorab instead.
        from entity.authorab_entity import AuthorabEntity
        return AuthorabEntity(self, data)


    @property
    def combined_search(self):
        """Idiomatic facade: client.combined_search.list() / client.combined_search.load({"id": ...})."""
        from entity.combined_search_entity import CombinedSearchEntity
        cached = getattr(self, "_combined_search", None)
        if cached is None:
            cached = CombinedSearchEntity(self, None)
            self._combined_search = cached
        return cached

    def CombinedSearch(self, data=None):
        # Deprecated: use client.combined_search instead.
        from entity.combined_search_entity import CombinedSearchEntity
        return CombinedSearchEntity(self, data)


    @property
    def combined_search_with_field(self):
        """Idiomatic facade: client.combined_search_with_field.list() / client.combined_search_with_field.load({"id": ...})."""
        from entity.combined_search_with_field_entity import CombinedSearchWithFieldEntity
        cached = getattr(self, "_combined_search_with_field", None)
        if cached is None:
            cached = CombinedSearchWithFieldEntity(self, None)
            self._combined_search_with_field = cached
        return cached

    def CombinedSearchWithField(self, data=None):
        # Deprecated: use client.combined_search_with_field instead.
        from entity.combined_search_with_field_entity import CombinedSearchWithFieldEntity
        return CombinedSearchWithFieldEntity(self, data)


    @property
    def line(self):
        """Idiomatic facade: client.line.list() / client.line.load({"id": ...})."""
        from entity.line_entity import LineEntity
        cached = getattr(self, "_line", None)
        if cached is None:
            cached = LineEntity(self, None)
            self._line = cached
        return cached

    def Line(self, data=None):
        # Deprecated: use client.line instead.
        from entity.line_entity import LineEntity
        return LineEntity(self, data)


    @property
    def linecount(self):
        """Idiomatic facade: client.linecount.list() / client.linecount.load({"id": ...})."""
        from entity.linecount_entity import LinecountEntity
        cached = getattr(self, "_linecount", None)
        if cached is None:
            cached = LinecountEntity(self, None)
            self._linecount = cached
        return cached

    def Linecount(self, data=None):
        # Deprecated: use client.linecount instead.
        from entity.linecount_entity import LinecountEntity
        return LinecountEntity(self, data)


    @property
    def poemcount(self):
        """Idiomatic facade: client.poemcount.list() / client.poemcount.load({"id": ...})."""
        from entity.poemcount_entity import PoemcountEntity
        cached = getattr(self, "_poemcount", None)
        if cached is None:
            cached = PoemcountEntity(self, None)
            self._poemcount = cached
        return cached

    def Poemcount(self, data=None):
        # Deprecated: use client.poemcount instead.
        from entity.poemcount_entity import PoemcountEntity
        return PoemcountEntity(self, data)


    @property
    def random(self):
        """Idiomatic facade: client.random.list() / client.random.load({"id": ...})."""
        from entity.random_entity import RandomEntity
        cached = getattr(self, "_random", None)
        if cached is None:
            cached = RandomEntity(self, None)
            self._random = cached
        return cached

    def Random(self, data=None):
        # Deprecated: use client.random instead.
        from entity.random_entity import RandomEntity
        return RandomEntity(self, data)


    @property
    def title(self):
        """Idiomatic facade: client.title.list() / client.title.load({"id": ...})."""
        from entity.title_entity import TitleEntity
        cached = getattr(self, "_title", None)
        if cached is None:
            cached = TitleEntity(self, None)
            self._title = cached
        return cached

    def Title(self, data=None):
        # Deprecated: use client.title instead.
        from entity.title_entity import TitleEntity
        return TitleEntity(self, data)


    @property
    def titleab(self):
        """Idiomatic facade: client.titleab.list() / client.titleab.load({"id": ...})."""
        from entity.titleab_entity import TitleabEntity
        cached = getattr(self, "_titleab", None)
        if cached is None:
            cached = TitleabEntity(self, None)
            self._titleab = cached
        return cached

    def Titleab(self, data=None):
        # Deprecated: use client.titleab instead.
        from entity.titleab_entity import TitleabEntity
        return TitleabEntity(self, data)



    @classmethod
    def test(cls, testopts=None, sdkopts=None):
        if sdkopts is None:
            sdkopts = {}
        sdkopts = vs.clone(sdkopts)
        if not isinstance(sdkopts, dict):
            sdkopts = {}

        if testopts is None:
            testopts = {}
        testopts = vs.clone(testopts)
        if not isinstance(testopts, dict):
            testopts = {}
        testopts["active"] = True

        vs.setpath(sdkopts, "feature.test", testopts)

        sdk = cls(sdkopts)
        sdk.mode = "test"

        return sdk
